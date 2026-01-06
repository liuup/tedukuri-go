import subprocess
import sys
import random

if len(sys.argv) < 2:
    print("Usage: python3 test_compare.py <basename>")
    sys.exit(1)

basename = sys.argv[1]
go_file = f"{basename}.go"
cpp_file = f"{basename}.cpp"
go_exec = f"{basename}_go"
cpp_exec = f"{basename}_cpp"

# Compile Go
subprocess.run(["go", "build", "-o", go_exec, go_file], check=True)

# Compile C++
subprocess.run(["g++", "-std=c++17", "-O2", "-o", cpp_exec, cpp_file], check=True)

def generate_input():
    chunks = []
    for _ in range(10):
        m = random.randint(2, 1000)
        n = random.randint(2, 1000)
        perm = list(range(m * n))
        random.shuffle(perm)

        # Build lines efficiently
        grid_lines = [" ".join(map(str, perm[i*n:(i+1)*n])) for i in range(m)]
        chunks.append(f"{m} {n}\n" + "\n".join(grid_lines))

    chunks.append("0 0")
    return "\n".join(chunks)


def run_program(cmd, input_str):
    proc = subprocess.run(cmd, input=input_str, text=True, capture_output=True)
    return proc.stdout.strip().rstrip("\n")  # remove trailing spaces/newlines

for i in range(10):
    test_input = generate_input()
    
    out_go = run_program([f"./{go_exec}"], test_input)
    out_cpp = run_program([f"./{cpp_exec}"], test_input)
    
    # Compare ignoring trailing spaces and at most one extra empty line
    go_lines = [line.rstrip() for line in out_go.splitlines()]
    cpp_lines = [line.rstrip() for line in out_cpp.splitlines()]
    
    # Allow one extra blank line at the end
    while go_lines and go_lines[-1] == "":
        go_lines.pop()
    while cpp_lines and cpp_lines[-1] == "":
        cpp_lines.pop()
    
    if go_lines != cpp_lines:
        print("INCORRECT FOR:")
        print(test_input)
        print("STANDARD:")
        print(out_cpp)
        print("YOUR PROGRAM:")
        print(out_go)
        print("Test case", i)
        sys.exit(0)
    else:
        print("Case", i, "passed")

print("ACCEPTED")

