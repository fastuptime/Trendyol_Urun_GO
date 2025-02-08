# 🛍️ Trendyol Ürün Bilgisi Çekici - GO 🚀

Bu proje, Trendyol API'sini kullanarak belirli bir ürün adına göre ürün bilgilerini çeken ve bu bilgileri JSON, CSV formatında kaydeden, ayrıca ürün resimlerini indiren bir Go uygulamasıdır.

---

## 📋 Özellikler

- **🔍 Ürün Arama**: Trendyol API'si üzerinden ürün arama.
- **📄 JSON & CSV Çıktısı**: Bulunan ürünleri JSON ve CSV formatında kaydetme.
- **🖼️ Resim İndirme**: Ürün resimlerini otomatik olarak indirme.
- **🎯 Maksimum Ürün Sınırı**: Kullanıcı tarafından belirlenen maksimum ürün sayısı.
- **📊 Terminal Arayüzü**: Kullanıcı dostu terminal arayüzü ile ilerleme takibi.

---

## 🛠️ Kurulum

1. **Go'yu Kurun**: Eğer sisteminizde Go yüklü değilse, [Go'yu indirip kurun](https://golang.org/dl/).

2. **Projeyi Klonlayın**:
   ```bash
   git clone https://github.com/fastuptime/Trendyol_Urun_GO.git
   cd Trendyol_Urun_GO
   ```

3. **Bağımlılıkları Kurun**:
   ```bash
   go mod download
   ```

4. **Projeyi Çalıştırın**:
   ```bash
   go run main.go
   ```

---

## 🖥️ Kullanım

1. **Ürün Adı Girin**: Program çalıştığında, aramak istediğiniz ürün adını girin.
2. **Maksimum Ürün Sayısı**: Kaç adet ürün bulmak istediğinizi belirtin (0 girerseniz sınırsız ürün aranır).
3. **Sonuçlar**: Bulunan ürünler `all_products.json` ve `all_products.csv` dosyalarına kaydedilir.
4. **Resimler**: Ürün resimleri `imgs` klasörüne indirilir.

---

## 📂 Dosya Yapısı

- `main.go`: Ana uygulama dosyası.
- `all_products.json`: Bulunan ürünlerin JSON formatında kaydedildiği dosya.
- `all_products.csv`: Bulunan ürünlerin CSV formatında kaydedildiği dosya.
- `imgs/`: İndirilen ürün resimlerinin bulunduğu klasör.

---

## 📸 Ekran Görüntüsü

![image](https://github.com/user-attachments/assets/b6b78c94-38c5-4d09-962c-398f3dee70dd)
![image](https://github.com/user-attachments/assets/7ef7e1d7-480b-41e9-9c62-79aeb629ba74)
![image](https://github.com/user-attachments/assets/3491314d-d15d-4911-a560-c0ecd8868952)
![image](https://github.com/user-attachments/assets/f6caa9c8-6601-465d-9680-bb1ae640059e)

---

## 🤖 Katkıda Bulunma

Katkıda bulunmak isterseniz, lütfen bir **Pull Request** gönderin. Yeni özellikler, hata düzeltmeleri veya dokümantasyon güncellemeleri her zaman memnuniyetle karşılanır!

---

## 💡 İletişim

Herhangi bir sorunuz veya öneriniz varsa, [GitHub Issues](https://github.com/fastuptime/Trendyol_Urun_GO/issues) üzerinden bize ulaşabilirsiniz.

---

**🌟 Star vererek projeyi destekleyebilirsiniz!**  
**🔗 Paylaşarak daha fazla kişiye ulaşmasını sağlayabilirsiniz!**
