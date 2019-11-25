#####################
Message State Diagram
#####################

.. versionadded:: v2.0.0
.. note::
   Only applicable to `v2.0+ Golang port <https://github.com/RITlug/teleirc/issues/163>`_.

This document explains possible states of a Message as it travels through TeleIRC.
There are two possible pathways, depending if a message is sourced from Telegram or IRC:

.. image:: /_static/dev/message-state-diagram.png
   :alt: TeleIRC Message State Diagram: raw message -> update object -> message object -> string -> pass message to other platform
