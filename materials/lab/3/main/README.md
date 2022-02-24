Aram Maljanian

I extended main.go to use my dns resolver functionality 
by first switching the program over to a flag based system
instead of using os.Args. The available flags are as follows:

-dnsresolve=uwyo.edu    Input web address, ip will be returned
-showcredits            If this flag is in the command the credits will be shown
-hosttosearch=Laramie   This was the original host search option

An example usage with fully featured output:
SHODAN_API_KEY=YOUR_API_KEY ./main -showcredits -hosttosearch=laramie -dnsresolve=uwyo.edu

Run with minimal options to just resolve a DNS address:
SHODAN_API_KEY=JZ0k75a7LqCOOmVBcbzozYwtj3IDLFLj ./main -dnsresolve=google.com

To resolve multiple addresses (separate by commas):
SHODAN_API_KEY=JZ0k75a7LqCOOmVBcbzozYwtj3IDLFLj ./main -dnsresolve=google.com,uwyo.edu,facebook.com