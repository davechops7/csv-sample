# csv-sample
Example demos showing .csv generation via GO.

After some research, reading and playing around with code I think the GO standard library gives us what we need for CSV Report tasks purpose. 

“The standard library is very robust. Consider standard library packages over bulky third-party frameworks” https://medium.com/@KeithAlpichi/gos-standard-library-by-example-encoding-csv-75f098169822
 
### Some alternatives:

[csvutil](https://github.com/jszwec/csvutil):
- Based on the Reader and Writer interfaces which are implemented by std Go csv package. 
 Offers best encoding and decoding performance of the 3.

[gocsv](https://github.com/gocarina/gocsv):
- Based on the Reader and Writer interfaces which are implemented by std Go csv package.
Customisable csv Reader/Writer.

[easycsv](https://github.com/yunabe/easycsv):
- Read focused only.
