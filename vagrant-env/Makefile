export HOST_NAME=publisher
export HOST_IP=192.168.34.11

.PHONY: up
up:
	vagrant up

.PHONY: halt
halt:
	vagrant halt

.PHONY: status
status:
	vagrant global-status

.PHONY: ssh
ssh:
	sshrc \
		-i .vagrant/machines/${HOST_NAME}/virtualbox/private_key \
		vagrant@${HOST_IP}

# デバッグ用
.PHONY: sync
sync:
	rsync \
		--recursive \
		--times \
		--compress \
		--verbose \
		-e 'ssh -i .vagrant/machines/${HOST_NAME}/virtualbox/private_key' \
		./ansible vagrant@${HOST_IP}:/home/vagrant/
