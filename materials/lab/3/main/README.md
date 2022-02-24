Aram Maljanian

I extended main.go to use my dns resolver functionality 
by first switching the program over to a flag based system
instead of using os.Args. The available flags are as follows:

-dnsresolve=uwyo.edu    Input web address, ip will be returned
-showcredits            If this flag is in the command the credits will be shown
-hosttosearch=Laramie   This was the original host search option