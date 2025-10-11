from PIL import Image
import sys
from PIL.ExifTags import TAGS

def print_exif_data(exif_data):
    for tag_id in exif_data:
        tag = TAGS.get(tag_id, tag_id)
        content = exif_data.get(tag_id)
        print(f'{tag:25}: {content}')




def main():
    if len(sys.argv) < 2:
        print("Usage: ")
        print("python3 scorpion.py FILE1 [FILE2 ...]")
        return

    
    print("Metadata of images")
    for arg in sys.argv[1:]:
        try:
            with Image.open(arg) as im:
                exifdata = im.getexif()
                if len(exifdata) == 0 and len(exifdata.get_ifd(0x8769)) == 0:
                    continue

                print(f"--- {arg} ---")
                print_exif_data(exifdata)
                print_exif_data(exifdata.get_ifd(0x8769))
        except Exception as e:
            print("scorpion err ", e)

main()
