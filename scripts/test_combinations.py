import json
import os
import subprocess


def run_combos(board, hands) -> dict:
    bin_path = './cmd/go-poker-combinations/go-poker-combinations'
    assert os.path.exists(bin_path), 'have no combinations executable'
    return json.loads(subprocess.run([bin_path, board, ','.join(hands)], capture_output=True).stdout)


if __name__ == '__main__':
    print(run_combos('Ks7s2s', ['2h2c']))
