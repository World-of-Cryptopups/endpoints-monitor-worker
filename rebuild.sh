git pull
go build -o emw 
sudo systemctl restart emw.service
journalctl -f -u emw.service