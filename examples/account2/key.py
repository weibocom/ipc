from steembase.account import PasswordKey

account = 'initminer'
password = '5JzpcbsNCu6Hpad1TYmudH4rj1A22SW9Zhb1ofBGHRZSp5poqAX'
key_types = [
    'posting',
    'active',
    'owner',
    'memo'
]

for key_type in key_types:
    private_key = PasswordKey(account, password, key_type).get_private_key()
    public_key = private_key.pubkey

    print('Private ' + key_type + ' key: ' + str(private_key))
    print('Public ' + key_type + ' key: ' + str(public_key) + '\n')
