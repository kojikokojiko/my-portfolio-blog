# my-portfolio-blog




エラー

````
TASK [vps : Synchronize frontend directory excluding node_modules] ***********************************************
fatal: [my-portfolio -> localhost]: FAILED! => {"changed": false, "cmd": "/usr/bin/rsync --delay-updates -F --compress --archive --rsh='/usr/bin/ssh -S none -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null' --rsync-path='sudo -u root rsync' --no-motd --exclude=node_modules --out-format='<<CHANGED>>%i %n%L' /Users/koji.iwase/mydev/my-portfolio-blog/ansible/frontend/ ubuntu@tk2-412-47176.vs.sakura.ne.jp:/var/www/frontend/", "msg": "Warning: Permanently added 'tk2-412-47176.vs.sakura.ne.jp' (ED25519) to the list of known hosts.\r\nubuntu@tk2-412-47176.vs.sakura.ne.jp: Permission denied (publickey).\r\nrsync: connection unexpectedly closed (0 bytes received so far) [sender]\nrsync error: unexplained error (code 255) at /AppleInternal/Library/BuildRoots/ce725a5f-c761-11ee-a4ec-b6ef2fd8d87b/Library/Caches/com.apple.xbs/Sources/rsync/rsync/io.c(453) [sender=2.6.9]\n", "rc": 255}
````


rsync -avz --exclude='node_modules' -e 'ssh -i ~/.ssh/my-portfolio/my-portfolio' /path/to/local/frontend/ ubuntu@tk2-412-47176.vs.sakura.ne.jp:/var/www/frontend/


````
- name: Ensure the target directory exists
  file:
    path: /var/www/frontend
    state: directory
    owner: ubuntu 
    group: www-data
    mode: '0775'

````


````


TASK [vps : Synchronize frontend directory excluding node_modules] ***********************************************
fatal: [my-portfolio -> localhost]: FAILED! => {"changed": false, "cmd": "/usr/bin/rsync --delay-updates -F --compress --archive --rsh='/usr/bin/ssh -S none -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null' --rsync-path='sudo -u root rsync' --no-motd --exclude=node_modules --out-format='<<CHANGED>>%i %n%L' /Users/koji.iwase/mydev/my-portfolio-blog/ansible/../frontend/ my-portfolio:/var/www/frontend/", "msg": "ssh: Could not resolve hostname my-portfolio: nodename nor servname provided, or not known\r\nrsync: connection unexpectedly closed (0 bytes received so far) [sender]\nrsync error: unexplained error (code 255) at /AppleInternal/Library/BuildRoots/ce725a5f-c761-11ee-a4ec-b6ef2fd8d87b/Library/Caches/com.apple.xbs/Sources/rsync/rsync/io.c(453) [sender=2.6.9]\n", "rc": 255}

````