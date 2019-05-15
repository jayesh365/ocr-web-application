try:
    from PIL import Image
except ImportError:
    import Image
import pytesseract


def formOcr(imgPath):
    '''
    (string) -> string
    This function will take in an image at <imgPath> and
    attempt to parse the text in the image.
    '''
    # initialize variable to represent the parsed text
    # Pillow will be used to open the image, and PyTesseract
    # will be used to parse strings from the image
    parsedText = pytesseract.image_to_string(Image.open(imgPath))
    # put all parsed text into an array
    parsedTextArr = [i for i in parsedText]
    return parsedText
