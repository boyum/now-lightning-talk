from http.server import BaseHTTPRequestHandler
from datetime import datetime


class handler(BaseHTTPRequestHandler):

    def do_GET(self):
        self.send_response(200)
        self.send_header('Content-type', 'application/json')
        self.end_headers()
        self.wfile.write(
            """
      {
        "time": "{time}",
        "lang": "python"
      }
      """.format(time=str(datetime.now().strftime('%Y-%m-%d %H:%M:%S')).encode()))
        return
