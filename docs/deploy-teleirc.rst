#####################
How to deploy TeleIRC
#####################

There are several ways to deploy TeleIRC persistently.
This page offers suggestions on possible deployment options.


*******
systemd
*******

The **recommended deployment method** is with `systemd <https://en.wikipedia.org/wiki/Systemd>`_.
This method requires a basic understanding of systemd unit files.
Create a unique systemd service for each TeleIRC instance.

A `provided systemd service <https://github.com/RITlug/teleirc/blob/master/misc/teleirc.service>`_ file is available.
Add the systemd unit file to ``/usr/lib/systemd/system/`` to activate it.
Now, ``systemctl`` can be used to control your TeleIRC instance.

Note the provided file makes two assumptions:

- Using a dedicated system user (e.g. ``teleirc``)
- TeleIRC config files located at ``/usr/lib64/teleirc/`` (i.e. files inside TeleIRC repository)


***
pm2
***

`pm2 <http://pm2.keymetrics.io/>`_ keeps NodeJS running in the background.
If you run an application and it crashes, pm2 restarts the process.
pm2 also restarts processes if the server reboots.
Read the `pm2 documentation <http://pm2.keymetrics.io/docs/usage/quick-start/>`_ for more information.

After pm2 is installed, run these commands to start TeleIRC::

    cd teleirc/
    pm2 start -n my-teleirc-bot teleirc.js


**************
Arch Linux AUR
**************

On ArchLinux, see `teleirc-git <https://aur.archlinux.org/packages/teleirc-git/>`_ in the AUR.
The AUR package uses the systemd method to deploy TeleIRC.
Place TeleIRC configuration files in the ``/usr/lib/teleirc/`` directory.


******
Docker
******

Docker is another way to deploy TeleIRC.
Dockerfiles and images are available in ``images/``.

Which image do I choose?
========================

Node Alpine Linux and Fedora images are provided.

+------------------------------------------------------------------------------+-----------------------+---------+
| Image                                                                        | File                  | Size    |
+==============================================================================+=======================+=========+
| `Node Alpine Linux <https://hub.docker.com/r/_/node/>`_ (``node:10-alpine``) | ``Dockerfile.alpine`` | 374 MB  |
+------------------------------------------------------------------------------+-----------------------+---------+
| `Fedora latest <https://hub.docker.com/r/_/fedora/>`_                        | ``Dockerfile.fedora`` | 569 MB  |
+------------------------------------------------------------------------------+-----------------------+---------+

This guide uses ``alpine``.
If you wish to use ``fedora``, replace ``alpine`` with ``fedora``.

Building Docker image
=====================

You may see errors running ``yarn``.
You can safely ignore them.
They are not fatal.

.. code-block:: bash

    docker build . -f images/Dockerfile.alpine -t teleirc
    docker run -d -u teleirc --name teleirc --restart always \
        -e TELEIRC_TOKEN="000000000:AAAAAAaAAa2AaAAaoAAAA-a_aaAAaAaaaAA" \
        -e IRC_CHANNEL="#channel" \
        -e IRC_BOT_NAME="teleirc" \
        -e IRC_BLACKLIST="CowSayBot,AnotherNickToIgnore" \
        -e TELEGRAM_CHAT_ID="-0000000000000" \
        teleirc

Docker Compose
==============

Optionally, you may use `docker-compose <https://docs.docker.com/compose>`_.
An `example compose file <https://github.com/RITlug/teleirc/blob/master/images/docker-compose.yml.example>`_ is provided.

Run these commands to use Docker Compose:

.. code-block:: bash

    cp images/docker-compose.yml.example docker-compose.yml
    cp env.example .env
    docker-compose up -d teleirc
