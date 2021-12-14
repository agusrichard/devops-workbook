from generate_table import initialize_db
from generate_table.create_table import generate_recipes

if __name__ == '__main__':
    ddb = initialize_db()
    generate_recipes(ddb)
