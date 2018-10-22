import os
import subprocess
import sys

def Boom(message):
    sys.stderr.write(":-( {0}\n".format(message))
    sys.exit(1)

def ImageExists(name):
    with open('/dev/null', 'w') as dev_null:
        return subprocess.call(
            ['docker', 'image', 'inspect', name],
            stdout=dev_null,
            stderr=dev_null) == 0

def BuildImage(name, dockerfile, context):
    return subprocess.call([
        'docker',
        'build',
        '--tag', name,
        '--file', dockerfile,
        context
    ]) == 0