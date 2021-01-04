# strong-pw-generator
Use GO to generate a strong password

Generates a password based on configurable parameters. Combining lower/upper/special character sets and numbers. 

It will automatically copy the generated password to your clipboard for easy pasting into pwpush. 

TODO:
- [x] CLI Accepts gen-pw cmd
- [x] Stores generated password in local redis
- [x] Generates UUID
- Stores UUID as key for redis
- Generate default TTL 
- Accept user defined TTL
- Research local pw push fork for storing/pushing passwords or roll my own