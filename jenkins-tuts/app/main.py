import fire


def hello(name="Veni"):
    return f"Hello {name}!"


if __name__ == "__main__":
    fire.Fire(hello)
