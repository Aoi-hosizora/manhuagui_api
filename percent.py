import os
from os import path
import re

log_filenames = [f for f in os.listdir('.') if f.endswith('.log') and f != 'console.current.log']
latencies = []

for log_filename in log_filenames:
    with open(log_filename, 'r', encoding='utf-8') as f:
        content = f.read()
        extracted = re.findall(r'"latency":(\d+),', content)
        extracted = [int(l) / 1e9 for l in extracted if int(l) < 8e9]
        latencies.extend(extracted)

print("Length of latencies:", len(latencies))

latencies = sorted(latencies)
for i in range(90, 100):
    p_percent = int(len(latencies) * i / 100)
    print("P{}: {}".format(i, latencies[p_percent - 1]))
