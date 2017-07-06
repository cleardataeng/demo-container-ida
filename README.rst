ida
===

``ida`` is a small and simple demo web service.  It uses `dactyl
<https://github.com/cleardataeng/demo-container-dactyl>`_ as an
example of composing services.

``ida`` supports a single endpoint at ``/``, and only handles ``GET``
requests.

A ``GET /`` response is a JSON data structure with the following
elements:

* ``Hostname``: the hostname where ``ida`` is running
* ``RemoteAddr``: the IP address from which ``ida`` received your
  request
* ``DactylInfo``: a JSON structure with the response from ``dactyl``
