############
Using Docker
############

Dockerfiles and images are available in ``containers/`` for configuring and running Teleirc.
Install Docker onto the machine you plan to run Teleirc from.


************************
Which image do I choose?
************************

Node Slim, Official Node Alpine Linux, and Fedora images are provided (ordered ascending by size).

+-----------------------------------------------------------------------------+-----------------------+---------+
| Image                                                                       | File                  | Size    |
+=============================================================================+=======================+=========+
| `Node Alpine Linux <https://hub.docker.com/r/_/node/>`_ (``node:6-alpine``) | ``Dockerfile.alpine`` | 66.9 MB |
+-----------------------------------------------------------------------------+-----------------------+---------+
| `Node Slim <https://hub.docker.com/r/_/node/>`_  (``node:6-slim``)          | ``Dockerfile.slim``   | 256 MB  |
+-----------------------------------------------------------------------------+-----------------------+---------+
| `Fedora latest <https://hub.docker.com/r/_/fedora/>`_                       | ``Dockerfile.fedora`` | 569 MB  |
+-----------------------------------------------------------------------------+-----------------------+---------+

This guide uses ``alpine``.
If you use another image, replace ``alpine`` with ``slim`` or ``fedora``.

You will see errors during ``npm install``.
You can safely ignore them.
They are not fatal.


*********************
Building Docker image
*********************

.. code-block:: bash

    cd containers/
    docker build . -f Dockerfile.alpine -t teleirc
    docker run -d -u teleirc --name teleirc --restart always \
        -e TELEIRC_TOKEN="000000000:AAAAAAaAAa2AaAAaoAAAA-a_aaAAaAaaaAA" \
        -e IRC_CHANNEL="#channel" \
        -e IRC_BOT_NAME="teleirc" \
        -e IRC_BLACKLIST="CowSayBot,AnotherNickToIgnore" \
        -e TELEGRAM_CHAT_ID="-0000000000000" \
        teleirc


**************
Docker Compose
**************

Optionally, you may use `docker-compose <https://docs.docker.com/compose>`_.
We provide an example compose file (``containers/docker-compose.yml.example``).

.. code-block:: yaml

    version: '2'
    services:
      teleirc:
        user: teleirc
        build:
          context: .
          dockerfile: Dockerfile.alpine
        env_file: .env

We ignore the ``docker-compose.yml`` file in ``.gitignore``.

Running with Compose
====================

Run these commands to begin using Teleirc with Docker Compose.

#. Copy ``containers/docker-compose.yml.example`` to ``docker-compose.yml``
#. ``docker-compose up -d teleirc``
