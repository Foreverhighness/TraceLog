#!/usr/bin/env python3
# -*- coding:utf-8 -*-
import time
import requests
import signal
import sys
import random


def getlines():
    ret = []
    with open('test.txt', mode='r') as f:
        ret = f.readlines()
    return ret


def gettextByurl():
    header = {
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36',
    }
    url = 'http://whatthecommit.com/index.txt'
    html = requests.get(url, headers=header)
    return html.text


def main():
    def exit(signum, frame):
        print('You choose to stop me.')
        sys.exit()
    signal.signal(signal.SIGINT, exit)
    signal.signal(signal.SIGTERM, exit)

    start, end = 1, 3
    lines = getlines()

    def gettext():
        return lines[random.randint(0, len(lines) - 1)]

    while True:
        i = random.randint(start, end)
        filename = f'{i}.log'
        with open(filename, mode='a') as f:
            text = f'{filename}: {gettext()}'
            print(text, end='')
            f.writelines(text)
        time.sleep(1)


if __name__ == '__main__':
    main()
