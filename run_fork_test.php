<?php

require_once 'vendor/autoload.php';

function issue_rpc(
    string $host,
    array $items
) : array {
    printf("%s\n", __FUNCTION__);
    $client = new Pkg\MaffClient(
        $host,
        ['credentials' => Grpc\ChannelCredentials::createInsecure()]
    );

    $req = new \Pkg\ProductReq();
    $req->setItems($items);
    printf("sending rpc\n");
    list ($res, $data) = $client->GetProduct($req)->wait();
    printf("sent rpc\n");

    // if (!$res) {
    //     throw new Exception(var_export($data, true));
    // }

    return [
        'product' => $res->getProduct(),
        'duration' => $res->getDuration(),
        'pid' => getmypid(),
    ];
}

function send_json(array $data) {
    printf("%s\n", json_encode($data));
}

$rpc_host = 'localhost:9090';
$should_fork = true;
$responses = [];
$items = [10, 20];

if ($should_fork) {
    $from_parent = issue_rpc($rpc_host, $items);
    $pid = pcntl_fork();
    if ($pid == -1) {
        die('unable to forkk');
    } elseif ($pid) {
        // $from_parent_after_fork = issue_rpc($rpc_host, $items);
        // send_json($from_parent_after_fork);
        $status = null;
        pcntl_waitpid($pid, $status);
        printf("parent is done\n");
    } else {
        printf("child: %s\n", getmypid());
        send_json([
            'from_parent' => $from_parent,
            'from_child' => issue_rpc($rpc_host, $items),
        ]);
        printf("child is done\n");
    }
} else {
    send_json(issue_rpc($rpc_host, $items));
}