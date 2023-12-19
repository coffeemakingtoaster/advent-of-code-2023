import os
import sys
import shutil
import time
import subprocess

from collections import defaultdict

PARTICIPANTS = ["coffeemakingtoaster"]

def get_dirs(path: str) -> list[str]:
    return [d for d in os.listdir(path) if os.path.isdir(os.path.join(path, d))]

def benchmark_folder(name: str) -> float:
    # Copy input files
    solution_dirs = get_dirs(f"./{name}")
    input_dirs = get_dirs("./data")
    timer_sum = 0
    
    for d in input_dirs:
        if d not in solution_dirs:
            print(f"Incomplete {d}")
            continue
        input_options = get_dirs(os.path.join("data",d))
        for option in input_options:
            shutil.copy(os.path.join("data", d, option, "input.txt"), os.path.join(name, d, "input.txt"))
            expected_res = ""
            with open(os.path.join("data", d, option, "output.txt")) as f:
                expected_res = f.read()
                expected_res = expected_res.replace("\n", "")
            start_time = time.time()
            res = subprocess.run("./main", cwd=os.path.join(name, d),shell=True, check=True, capture_output=True, text=True)
            timer_sum += (time.time() - start_time)
            output = res.stdout
            output = output.replace("\n","")
            if output != expected_res:
                print(f"{output}:{expected_res}")
                sys.exit(1)
            os.remove(os.path.join(name, d, "input.txt"))
            
    return timer_sum
    

def main() -> None:
    benchmarks = defaultdict(lambda: 0)
    print(f"Benchmarking {len(PARTICIPANTS)} participants")
    for participant in PARTICIPANTS:
        benchmarks[participant] = benchmark_folder(participant)
        print(f"{participant}:{benchmarks[participant]}")
    

if __name__ == "__main__":
    main()