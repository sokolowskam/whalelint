#STS005
FROM some_very_trusted/ubuntu:18.04

#RUN013
RUN apt-get install -y --no-install-recommends gnupg=2.2.19 software-properties-common=0.98 apt-utils=2.0.9 openssh-client=8.2

#RUN011
RUN ssh-keygen -q -t rsa -N "" -f /id_rsa

#RUN012
RUN useradd -ms /bin/bash ubuntu && \
    usermod -aG sudo ubuntu
RUN echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

USER ubuntu
RUN sudo mkdir -p /opt/packages/gradle
WORKDIR /opt/packages/gradle

#RUN014, RUN016
RUN sudo wget --user=admin --password=admin123 https://services.gradle.org/distributions/gradle-3.1-bin.zip;  \
      sudo unzip gradle-3.1-bin.zip;  \
      sudo ln -s /opt/packages/gradle/gradle-3.1/bin/gradle /usr/bin/gradle

WORKDIR /app

#ADD001
ADD config config/

#CPY007
COPY .. .

USER root

EXPOSE 22
CMD ["/usr/sbin/sshd", "-D"]

