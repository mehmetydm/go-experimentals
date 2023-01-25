import re
with open("wordlist.txt", "r") as f:
    lines = f.readlines()
pattern = re.compile(r'::(.+?)::(.+)')
prepare = {}
for line in lines:
    match = pattern.match(line)
    if match:
        print(line)
        print(line[0:line.rfind('::')] )
        print(match.group(1))
        line = line[0:line.rfind('::')] + '::' + match.group(1)
        prepare[match.group(1)] = match.group(2)

