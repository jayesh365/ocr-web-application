from pdf2image import convert_from_path


def pdfToImage(pdf):
    '''
    (string) -> null
    this function takes in a path to a pdf
    and converts the pdf intp an image (type jpeg).
    '''

    # convert pages into images
    pages = convert_from_path(pdf)

    # save the pages in jpeg format
    for page in pages:
        page.save('./uploads/page.jpg', 'JPEG')
