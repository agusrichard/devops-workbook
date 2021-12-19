from generate_table import initialize_db
from generate_table.recipes import generate_recipes, drop_recipes

def generate_table():
    ddb = initialize_db()
    generate_recipes(ddb)

def drop_table():
    ddb = initialize_db()
    drop_recipes(ddb)

if __name__ == '__main__':
    generate_table()
