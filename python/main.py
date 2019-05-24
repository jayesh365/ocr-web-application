from ocr import formOcr
from pageConverter import crop


def main():
    image = "./uploads/form.tiff"
    crop(image, (269, 908, 2259, 1022), './uploads/date.jpg')
    crop(image, (270, 1110, 2260, 1435), './uploads/comments.jpg')
    print(formOcr('./uploads/date.jpg'))
    print(formOcr('./uploads/comments.jpg'))


main()
