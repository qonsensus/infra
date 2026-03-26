# Quonsensus Infrastructure Installation helper (docker version)
# This script shall be selfcontained for easy curl | python installation
# You may only use pythons standard librarys

import os
import sys
from dataclasses import dataclass
from subprocess import run
from time import sleep


# Snaity Checks
@dataclass
class Env:
    python_version: float = 3.7
    docker_version: str = ""


def detectEnv():
    return Env(
        python_version=float(f"{sys.version_info[0]}.{sys.version_info[1]}"),
        docker_version=run(["docker", "version"], capture_output=True, text=True)
        .stdout.strip()
        .split()[2]
        .strip(","),
    )


# Utils
class Colors:
    HEADER = "\033[95m"
    OKBLUE = "\033[94m"
    OKCYAN = "\033[96m"
    OKGREEN = "\033[92m"
    WARNING = "\033[93m"
    FAIL = "\033[91m"
    ENDC = "\033[0m"
    BOLD = "\033[1m"
    UNDERLINE = "\033[4m"


def user_input_str(question: str, default: str = "") -> str:
    if default == "":
        return input(f"{question}:")
    else:
        message = input(f"{question}  [{default}]:")

    if not message and default != "":
        return default
    elif not message and default == "":
        return user_input_str(question, default)

    return message


def user_input_int(question: str, default: int = 0) -> int:
    if default == 0:
        return int(input(f"{question}:"))
    else:
        message = input(f"{question}  [{default}]:")

    try:
        message = int(message)
    except ValueError:
        print(f"{Colors.FAIL}Invalid input. Please enter a number.{Colors.ENDC}")
        return user_input_int(question, default)

    return int(message)


def user_input_bool(question: str) -> bool:
    yes_choices = ["yes", "y"]
    no_choices = ["no", "n"]

    message = input(question + " (y / n): ")
    message = message.strip().lower()

    if message in yes_choices:
        return True
    elif message in no_choices:
        return False
    else:
        return user_input_bool("Please answer with y (yes) or n (no): ")


def rmLine():
    print("\033[F\033[K", end="")


# Configuration Data Structures
@dataclass
class Database:
    DB_HOST: str = "localhost"
    DB_PORT: int = 5432
    DB_USER: str = "admin"
    DB_PASSWORD: str = "better be secure!"
    DB_NAME: str = "quonsensus"
    DB_EXPOSE: bool = False


@dataclass
class Backend:
    BE_EXPOSE: bool = False
    BE_PORT: int = 7000


@dataclass
class Frontend:
    FE_EXPOSE: bool = False
    FE_PORT: int = 7001


if __name__ == "__main__":
    print(f"Welcome to {Colors.BOLD}Quonsensus{Colors.ENDC}")
    env = detectEnv()
    if env.python_version > 3.7:
        print("Python version 3.7 or higher is required for this installation.")
        sys.exit(1)

    if not env.docker_version:
        print("Docker is required to install the qonsensus")
        sys.exit(1)
