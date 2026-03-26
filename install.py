# Quonsensus Infrastructure Installation helper (docker version)
# This script shall be selfcontained for easy curl | python installation
# You may only use pythons standard librarys

import string
import sys
from dataclasses import dataclass
from random import choice
from subprocess import run


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


def rm_line():  # TODO: Fix added newlines and properly work into user_input functions
    print("\033[F\033[K", end="")


def user_input_str(question: str, default: str = "") -> str:
    if default == "":
        message = input(f"{question}:")
    else:
        message = input(f"{question}  [{default}]:")

    if not message and default != "":
        return default
    elif not message and default == "":
        rm_line()
        return user_input_str(question, default)

    return message


def user_input_int(question: str, default: int = 0) -> int:
    if default == 0:
        message = int(input(f"{question}:"))
    else:
        message = input(f"{question}  [{default}]:")

    if not message:
        return default

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
        rm_line()
        return user_input_bool("Please answer with y (yes) or n (no): ")


def generatePassword(length=20) -> str:
    alphabet = string.ascii_letters + string.digits
    return "".join(choice(alphabet) for i in range(length))


# TODO: do this at all lmao
def create_docker_compose(db: Database, be: Backend, fe: Frontend):
    pass


# Configuration Data Structures
@dataclass
class Database:
    DB_HOST: str = "localhost"
    DB_PORT: int = 5432
    DB_USER: str = "admin"
    DB_PASSWORD: str = "secure"
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

    print("\nThis script will guide you through the installation of Quonsensus.")
    print(f"{Colors.HEADER}We will start by setting up the database.{Colors.ENDC}")

    db = Database()
    db.DB_HOST = user_input_str("Enter the database host", db.DB_HOST)
    rm_line()
    db.DB_PORT = user_input_int("Enter the database port", db.DB_PORT)
    rm_line()
    db.DB_USER = user_input_str("Enter the database username", db.DB_USER)
    rm_line()
    db.DB_PASSWORD = user_input_str("Enter the database password", generatePassword())
    rm_line()
    db.DB_NAME = user_input_str("Enter the database name", db.DB_NAME)
    rm_line()
    db.DB_EXPOSE = user_input_bool("Expose the database to the host? ")
    rm_line()
    rm_line()
    print(f"{Colors.HEADER}Now configuring the backend...{Colors.ENDC}")
    be = Backend()
    be.BE_EXPOSE = user_input_bool("Expose the backend to the host? ")
    rm_line()
    be.BE_PORT = user_input_int("Enter the backend port", be.BE_PORT)
    rm_line()
    rm_line()
    print(f"{Colors.HEADER}Now configuring the frontend...{Colors.ENDC}")
    fe = Frontend()
    fe.FE_EXPOSE = user_input_bool("Expose the frontend to the host? ")
    rm_line()
    fe.FE_PORT = user_input_int("Enter the frontend port", fe.FE_PORT)
    rm_line()

    if fe.FE_PORT == be.BE_PORT:  # TODO: Test more potential errors
        print("The frontend and backend are running on the same port!")
    else:
        print(
            f"{Colors.OKGREEN}Configuration complete. Creating Docker containers...{Colors.ENDC}"
        )
        create_docker_compose(db, be, fe)
