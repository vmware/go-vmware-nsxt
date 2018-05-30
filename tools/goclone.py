import re
import sys

def main():
    if len(sys.argv) != 6:
        print("Usage: %s <filename> <func to replicate> <new func name> <replace what> <replace with>" % sys.argv[0])
        sys.exit(0)

    filename = sys.argv[1]
    func = sys.argv[2]
    new_func = sys.argv[3]
    what = sys.argv[4]
    to = sys.argv[5]

    with open(filename, 'r') as f:
        lines = f.readlines()

    lines_before_func = []
    lines_after_func = []
    bunch = []
    func_started = False
    func_ended = False
    for line in lines:
        if func_ended:
            lines_after_func.append(line)
            continue

        if line[:2] == '/*' and not func_started:
            lines_before_func.extend(bunch)
            bunch = [line]
        else:
            bunch.append(line)

        if re.match("func .+ %s" % func, line):
            func_started = True

        if func_started and line[:1] == '}':
            func_ended = True


    output = lines_before_func
    output.extend(bunch)
    output.append('\n')
    for line in bunch:
        replaced = re.sub('\\b' + what + '\\b', to, line)
        output.append(re.sub('\\b' + func + '\\b', new_func, replaced))

    output.extend(lines_after_func)

    for line in output:
        print(line[:-1])

main()
