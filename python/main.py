from ocr import formOcr
from pageConverter import crop
from pdfConverter import pdfToImage
from parseCSV import parse


def main():
    image = "./uploads/form.tiff"
    crop(image, (95, 634, 1557, 708), './uploads/date.jpg')
    crop(image, (99, 776, 1557, 1008), './uploads/comments.jpg')
    print(formOcr('./uploads/date.jpg'))
    print(formOcr('./uploads/comments.jpg'))


main()
