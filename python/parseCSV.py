import csv
import json
import os


def parse(file):
    # open csv file to parse
    with open(file, 'r') as csvFile:
        # read the file as a dict
        reader = csv.DictReader(csvFile)
        dictCsv = None
        rowNum = 0
        # put each question/answer pair as a dict and create a json file
        for row in reader:
            dictCsv = dict(row)
            json.dump(dictCsv, open(
                "./uploads/parsedCSV" + str(rowNum) + ".json", 'w'), indent=4, sort_keys=True)
            rowNum += 1
    # close the csv file
    csvFile.close()
    # delete csv file
    os.remove(file)


parse("./sdaps/data_1.csv")
