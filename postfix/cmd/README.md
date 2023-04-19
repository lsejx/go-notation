# Installation
	go install github.com/lsejx/go-notation/postfix/cmd/pfnc

# Usage

### help
	pfnc -h

### 64bit float mode
	pfnc 1 1 +         // 1+1=2
	pfnc pi 2 x        // 2π=6.28...
	pfnc 2 8 xx        // pow(2,8)=256
	pfnc 2 1 2 / xx    // sqrt(2)=1.41...
	pfnc 1.234 1k x    // 1234
	pfnc 1k 1m x       // 1

### 64bit signed int mode
	pfnc -i 1 1 +
	pfnc -i max
	pfnc -i min

### 64bit unsigned int mode
	pfnc -u 1 1 +
	pfnc -u max
