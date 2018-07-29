############
Using Docker
############

Dockerfiles and images are available in ``containers/`` for configuring and running Teleirc.
Install Docker onto the machine you plan to run Teleirc from.


************************
Which image do I choose?
************************

Official Node Alpine Linux, and Fedora images are provided (ordered ascending by size).

+-----------------------------------------------------------------------------+-----------------------+---------+
| Image                                                                       | File                  | Size    |
+=============================================================================+=======================+=========+
| `Node Alpine Linux <https://hub.docker.com/r/_/node/>`_ (``node:8-alpine``) | ``Dockerfile.alpine`` | 374 MB  |
+-----------------------------------------------------------------------------+-----------------------+---------+
| `Fedora latest <https://hub.docker.com/r/_/fedora/>`_                       | ``Dockerfile.fedora`` | 569 MB  |
+-----------------------------------------------------------------------------+-----------------------+---------+

This guide uses ``alpine``.
If you wish to use ``fedora``, replace ``alpine`` with ``fedora``.

You will see errors during ``yarn``.
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
We provide an example compose file (``containers/docker-compose.yml``).

.. code-block:: yaml

    version: '2'
    services:
      teleirc:
        user: teleirc
        build:
          context: ..
          dockerfile: containers/Dockerfile.alpine
        env_file: .env

Running with Compose
====================

Run these commands to begin using Teleirc with Docker Compose.


#. Copy ``docker-compose.yml.example`` to ``docker-compose.yml`` and edit if you do not wish to use the alpine image
#. Copy ``.env.example`` to ``.env`` and edit accordingly.
#. ``docker-compose up -d teleirc``
