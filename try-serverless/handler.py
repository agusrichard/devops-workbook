import json


def run(event, context):
    body = {
        "message": "Go Serverless v3.0! Your function executed successfully! Sekardayu",
        "input": event,
    }

    return {"statusCode": 200, "body": json.dumps(body)}
