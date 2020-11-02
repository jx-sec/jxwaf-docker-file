FROM centos:7

RUN yum update -y && yum install -y epel-release
RUN yum install -y python-devel  python-pip git  gcc make mysql mysql-devel
RUN cd /opt && git clone https://github.com/jx-sec/jxwaf-mini-server.git server
RUN cd /opt/server && pip install  -r requirements.txt

CMD cd /opt/server \
    && python manage.py makemigrations \
    && python manage.py migrate \
    && python manage.py runserver 0.0.0.0:80