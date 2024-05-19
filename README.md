# torrent-webplayer

# iptv

#init_d service 

-> in file /system/etc/init

+ create irtv.rc

```
service irtv /data/tmp/myscript.sh
    user root
    group root
    disabled
    oneshot

on property:sys.boot_completed=1 && property:sys.logbootcomplete=1
    startirtv
```

in myscript.sh file

```
#!/system/bin/sh
ir-keytable -c -w filemap.toml
```
