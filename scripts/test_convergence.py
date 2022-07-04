import json
import subprocess
import tqdm


opp_range = """
 9d9c:0.209,9h9c:0.212,9h9d:0.201,9s9c:0.209,9s9d:0.186,9s9h:0.201,AcJc:0.095,AcJd:0.095,AcJh:0.095,AcKc:0.044,AcKd:0.044,AcKh:0.056,AcKs:0.045,AcQc:0.045,AcQh:0.035,AcQs:0.046,AcTc:0.076,AcTd:0.077,AcTh:0.091,AcTs:0.077,AdAc:0.009,AdJc:0.095,AdJd:0.095,AdJh:0.095,AdKc:0.045,AdKd:0.045,AdKh:0.056,AdKs:0.045,AdQc:0.045,AdQh:0.035,AdQs:0.046,AdTc:0.076,AdTd:0.077,AdTh:0.091,AdTs:0.077,AhAc:0.015,AhAd:0.015,AhJc:0.1,AhJd:0.1,AhJh:0.2,AhKc:0.06,AhKd:0.06,AhKh:0.1,AhKs:0.06,AhQc:0.059,AhQh:0.097,AhQs:0.06,AhTc:0.104,AhTd:0.105,AhTh:0.1,AhTs:0.105,AsAc:0.009,AsAd:0.009,AsAh:0.015,AsJc:0.095,AsJd:0.095,AsJh:0.095,AsKc:0.045,AsKd:0.045,AsKh:0.056,AsKs:0.045,AsQc:0.045,AsQh:0.035,AsQs:0.046,AsTc:0.076,AsTd:0.077,AsTh:0.091,AsTs:0.077,JdJc:0.086,JhJc:0.068,JhJd:0.068,KcQc:0.04,KcQh:0.031,KcQs:0.04,KdKc:0.002,KdQc:0.04,KdQh:0.031,KdQs:0.04,KhKc:0.012,KhKd:0.012,KhQc:0.048,KhQh:0.08,KhQs:0.051,KsKc:0.002,KsKd:0.002,KsKh:0.012,KsQc:0.04,KsQh:0.031,KsQs:0.04,QhQc:0.065,Qs5c:0.274,QsQc:0.094,QsQh:0.063,TdTc:0.261,ThTc:0.301,ThTd:0.3,TsTc:0.261,TsTd:0.259,TsTh:0.301
 """.strip()
board = '7hAsQdJs8h'
hand = '2dQc'


def run_calculator(iterations) -> dict:
    p = subprocess.Popen(f"./go-poker-equity --iter {iterations} {board} {hand} {opp_range}", stdout=subprocess.PIPE, shell=True)
    return json.loads(p.communicate()[0].decode().strip())


if __name__ == '__main__':
    equity = []
    for i in tqdm.tqdm(range(1000000, 20000000, 1000000)):
        equity.append(run_calculator(i)['equity']['2dQc'])
    print(equity)

