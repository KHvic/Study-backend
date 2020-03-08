import pandas as pd
import xlrd

questionFile = "./greq.xlsx"
outputFile = "../../db/init-mysql.sql"

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
        value = value + addQuote(row["answer"]) + "," + str(row["type"]) + ")"
        values.append(value)
    sql = "INSERT INTO `question` (`category`, `sub_category`, `description`, `options`, `answer`, `type`) VALUES"
    sql += ",".join(values)
    sql += ";\n"
    return sql
    
def main():
    file1 = open(outputFile, "w") 
    
    types = ["TC1", "SE"]
    for type in types:
        file1.write(generateSQL(type).encode('ascii', 'ignore').decode('ascii'))
    file1.close()
    
if __name__== "__main__":
    main()
