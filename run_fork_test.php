<?php

require_once 'vendor/autoload.php';

function issue_rpc(
    string $host,
    array $items
) : array {
    $client = new Pkg\MaffClient(
        $host,
        ['credentials' => Grpc\ChannelCredentials::createInsecure()]
    );

    $req = new \Pkg\ProductReq();
    $req->setItems($items);
    list ($res, $data) = $client->GetProduct($req)->wait();

    return [
        'product' => $res->getProduct(),
        'duration' => $res->getDuration(),
        'pid' => getmypid(),
    ];
}

function send_json(array $data) {
    printf("%s\n", json_encode($data, JSON_PRETTY_PRINT));
}

$rpc_host = 'localhost:9090';
$should_fork = true;
$responses = [];
$items = [10, 20];

printf(
    "Test test issues a rpc call from the parent process, then\n" .
    "forks and issues another rpc from the child process. The\n" .
    "parent waits on the child to terminate. The test should\n" .
    "print a single JSON message and then quit properly.\n"
);

if ($should_fork) {
    $from_parent = issue_rpc($rpc_host, $items);
    $pid = pcntl_fork();
    if ($pid == -1) {
        die('unable to forkk');
    } elseif ($pid) {
        // send_json(issue_rpc($rpc_host, $items));
        $status = null;
        pcntl_waitpid($pid, $status);
    } else {
        send_json([
            'from_parent' => $from_parent,
            'from_child' => issue_rpc($rpc_host, $items),
        ]);
    }
} else {
    send_json(issue_rpc($rpc_host, $items));
}