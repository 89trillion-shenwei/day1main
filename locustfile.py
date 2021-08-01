from locust import HttpUser, between, task


class WebsiteUser(HttpUser):
    wait_time = between(5, 15)

    @task
    def index1(self):
        self.client.get("/FindSoldierByRaUnCv?Rarity=2&UnlockArena=3&Cvc=1000")

    @task
    def index2(self):
        self.client.get("/FindSoldierRaById?Id=14509")

    @task
    def index3(self):
        self.client.get("/FindSoldierCoById?Id=14509")

    @task
    def index4(self):
        self.client.get("/FindSoldierByCv?Cvc=1000")

    @task
    def index5(self):
        self.client.get("/FindSoldierByUn")
