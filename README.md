# Mastermind
Savaşan 22 Yarışmasının sisteminde mesajları yönlendirmek ve sistem durumunu tutmaktan sorumlu Yönetici servis

## Tasarım
Bu program, farklı gRPC bağlantıları açarak diğer servislerle konuşabilecektir. Resimdeki iki taraflı yayınlama (bidirectional streaming) ile bağlantılar açık tutulacaktır.

![image](https://user-images.githubusercontent.com/53450844/177382414-6e5f8ecc-e955-4d49-9c04-8818763de7a3.png)

Mastermind, yönetici servis olup diyagramdaki diğer servislerle haberleşecektir.
![Sistem Mimarisi](https://user-images.githubusercontent.com/53450844/181495511-c5dd93b5-4f80-4de2-9b1e-53146efb052d.png)


## Sistem Durumu
Bağlı olan servislerin sistemin hangi aşamada olduğundan haberdar olabilmeleri için, Mastermind bütün servislerle bu sistem durumu paylaşmak zorunda. Sistemi bir sonlu durum makinesi (finite-state machine) olarak modelleyerek, belirli kurallara göre hangi aşamada olması gerektiğini hesaplayıp, X geçiş fonksiyonu için Y durumuna geçileceğini elde edilebilir. 2022 Savaşan İHA Yarışması'nın isterlerini yerine getirmek için kullanılacak sistem durumları ve geçiş fonksiyonları aşağıdaki durum diyagramında gösterilir:
![Durum Makinesi](https://user-images.githubusercontent.com/53450844/181495372-cb0da3ee-6a44-41cb-b599-208a7b66471c.png)



### Endpointlar / RPCler

#### UpdateState(stream UpdateStateRequest) returns (stream UpdateStateResponse)
UpdateState bi-directional RPC üzerinden istemci servislerden durum geçişleri alırken, onlara yeni hesaplanan durumu da gönderecektir. Unary yerine bi-directional olmasının nedeni, durum güncellemesinin sadece gönderen servis değil, diğer servislerden de kaynaklı olabilmesinden dolayıdır. Örn.

Servis A bir durum güncellemesini gönderdiği zaman, Servis A, Servis B, ve Servis C'ye yeni durumu içeren bir UpdateStateResponse gönderilecek.


## Quickstart

```
git clone https://github.com/koustech/mastermind.git
cd mastermind
go run server/main.go -a 127.0.0.1:5678
```


