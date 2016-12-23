.DEFAULT_GOAL = create
.PHONY: create
create: depends build


.PHONY: depends
depends: 
	@ pip install -r requirements.txt

.PHONY: build
build: 
	@ export PATH="/pyinstaller:$$PATH" \
		&& pyinstaller \
	    	--exclude-module pycrypto \
	    	--exclude-module PyInstaller \
	    	-F dockerstart.py


.PHONY: shell
shell: 
	@ /bin/sh


.PHONY: python
python: 
	@ /usr/bin/env python