from os import walk
import shutil
from pathlib import Path


def is_txt(s):
    return '.txt' in s


base_addr = '.'
_, _, file_names = next(walk(base_addr))
poems = list(filter(is_txt, file_names))

# print(poems)

for p in poems:
    typ = p.split('-')[0]
    school = p.split('-')[1]
    author = p.split('-')[2]
    name = p.split('-')[3][:-4]

    # print(typ, school, author, name)

    dst_addr = base_addr+'/'+typ+'/'+school+'/'+author
    Path(dst_addr).mkdir(parents=True, exist_ok=True)
    shutil.move(base_addr+'/'+p, dst_addr+'/'+p)
