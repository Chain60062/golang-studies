class Test:
    def __init__(self, x):
        self.x = x

    @staticmethod
    def outer():
        test = Test(10)
        Test.inner1(test)  # test.x agora vale 5
        print(test.x)

        Test.inner2(test)  # Não modifica a referencia de test
        print(test.x)

        g = None
        Test.inner2(g)
        print(g is None)

    @staticmethod
    def inner1(test):
        test.x = 5

    @staticmethod
    def inner2(test):
        test = Test(200)


Test.outer()
