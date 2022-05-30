#!/bin/bash

echo 'export WCloud_Mesh_SubType="1"' >> /etc/profile
source /etc/profile

##export WCloud_Mesh_SubType="1"

echo "read WCloud_Mesh_SubType=$WCloud_Mesh_SubType"