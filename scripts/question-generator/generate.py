import pandas as pd
import xlrd
import os.path as path

questionFile = path.abspath("./greq.xlsx")
baseFile = path.abspath("../../db/init-mysql-master.sql")
outputFile = path.abspath("../../db/data/init-mysql.sql")

def addQuote(str):
    return "'" + str + "'"
def addQuote2(str):
    return '"' + str + '"'
    
def generateSQL(type):
    df = pd.read_excel(questionFile, type)
    df['description'] = df['description'].str.replace('\"','\\"',regex=True)
    values = []
    for _, row in df.iterrows():
        value = "(" + addQuote(row["category"]) + "," + addQuote(row["sub_category"]) + ","
        value = value + addQuote2(row["description"]) + "," + addQuote(row["options"]) + ","
        value = value + addQuote(row["answers"]) + "," + str(row["type"]) + ")"
        values.append(value)
    sql = "INSERT INTO `question` (`category`, `sub_category`, `description`, `options`, `answers`, `type`) VALUES"
    sql += ",".join(values)
    sql += ";\n"
    return sql
    
def main():
    base = open(baseFile, "r")
    baseContent = base.read()

    output = open(outputFile, "w+") 
    output.write(baseContent)
    types = ["TC1", "SE"]
    for type in types:
        output.write(generateSQL(type).encode('ascii', 'ignore').decode('ascii'))

    output.close()
    base.close()
    
if __name__== "__main__":
    main()
