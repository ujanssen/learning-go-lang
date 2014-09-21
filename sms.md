Simple Message Service
======================

Definitions:

- An author send messages to one or more recipients.
- A message consists of the message header, and the message body.

The message
-----------

The message header contains control information, including, an author's email address and one or more recipient addresses. Usually descriptive information is also added, such as a subject header field and a message submission date/time stamp. 

- an author's email address
- one or more recipient addresses
- a message submission date/time stamp
- a subject header field
- a message body

A message can have box tags

- draft or sent (author)
- inbox, trash (recipient)

A message can have a state tags

- read or unread (recipient)
- deleted

The service
-----------

The service defines an api.
It must support the creation and deletion of users.
And an user can sent and receve messages.

Ressources:

    /sms/users
    /sms/message
    
