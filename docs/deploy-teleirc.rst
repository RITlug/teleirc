#####################
How to deploy TeleIRC
#####################

.. warning:: This page is not yet fully updated for TeleIRC v2.0.0 or later.
   It is a work-in-progress.
   When the instructions here are "production-ready," this warning will be removed.
   Follow `RITlug/teleirc#193 <https://github.com/RITlug/teleirc/issues/193>`_ and `RITlug/teleirc#228 <https://github.com/RITlug/teleirc/issues/228>`_ to track progress.

There are two ways to deploy TeleIRC persistently:

#. Run Go binary
#. Run TeleIRC in a container


*************
Run Go binary
*************

This section explains how to configure and install TeleIRC as a simple executable binary.

Requirements
============

- git
- go (v1.13.x or later)

Install dependencies
====================

#. Clone the repository (``git clone https://github.com/RITlug/teleirc.git``)
#. Install dependencies (``go install``)

Configuration
=============

TeleIRC uses `godotenv <https://github.com/joho/godotenv>`_ to manage API keys and settings.
The config file is a ``.env`` file.
Copy the example file to a production file to get started (``cp env.example .env``).
Edit the ``.env`` file with your API keys and settings.

.. seealso::

   See :doc:`config-file-glossary` for detailed information.

Start bot
=========

.. note:: Work in progress.


***************
Run a container
***************

Containers are another way to deploy TeleIRC.
Dockerfiles and other deployment resources are available in ``deployments/``.

.. important:: The below instructions only apply to v1.x.x releases.
   Once a v2.x.x container image is created, these instructions will change.

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
