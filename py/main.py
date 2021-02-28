from fastapi import FastAPI
import uvicorn
import json

app = FastAPI()

import grpc
from py.grpc_f import user_pb2_grpc
from py.grpc_f import user_pb2


def object_to_dict(obj): return obj.__dict__


@app.get('/')
def home_page():
    return 'Hello World from Python!'


@app.get('/user/{username}')
def get_user(username: str = 'Peter'):
    print("Started gRPC service")
    # try:
    channel = grpc.insecure_channel('localhost:3000')
    stub = user_pb2_grpc.UserStub(channel)
    userRequest = user_pb2.UserRequest(username=username)
    userResponse = stub.Search(userRequest)
    return {'name': userResponse.username}
    # except Exception as e:
    #     print(e)
    #     return e


if __name__ == '__main__':
    uvicorn.run(app, port="5000")
