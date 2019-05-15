import csv
import json


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
                "./uploads/parsedCSV" + str(rowNum) + ".json", 'w'))
            rowNum += 1
    # close the csv file
    csvFile.close()
