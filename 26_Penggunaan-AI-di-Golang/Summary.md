# Summary

## (26) Penggunaan AI di Golang

### Why GO ?
- Dikembangkan oleh Google untuk pengembangan aplikasi server-side dan berbasis cloud
- Go adalah compiled language sehingga lebih cepat dibanding interpreted language (contoh : Python)
- Mendukung concurrency melalui fitur goroutines sehingga dapat menjalankan banyak task secara simultan tanpa membenani memori
- Memiliki standard libraries yang sangat kaya, didukung oleh komunitas yang kuat
- Termasuk libraries untuk mendukung pengembangan AI

### Library GO untuk Machine Learning
- GoLearn 
    - Bersifat open source
    - Batteries included machine learning library
    - Simplicity & customizability
    - Contoh penerapan : [GoLearn](https://github.com/sjwhitworth/golearn#getting-started)
- Gorgonia 
    - Memudahkan menulis dan mengevaluasi persamaan matematika yang melibatkan array multidimensi
    - Low-level library, seperti Theano, namun lebih tinggi dibanding Tensorflow
    - Contoh penerapan :  [Gorgonia](https://github.com/gorgonia/gorgonia#usage)
- Gonum
    - Sebuah set packages yang didesain untuk menulis numerical & algoritma sains menjadi produktif, memiliki performa tinggi, dan scalable
    - Mirip seperti numpy dan spicy, library yang dibuat menggunakan Python
    - Contoh penerapan : [Gonum](https://www.gonum.org/post/intro_to_gonum/)
- Gomind
    - Contoh penerapan : [Gomind](https://github.com/surenderthakran/gomind#usage)

### AI as a services (AIaaS)
- Sebuah tools AI yang dapat segera digunakan (pre-built or off-the-shelf AI tools)
- Berguna bagi bisnis yang ingin menerapkan AI tanpa merekrut ahlinya dan tanpa mengeluarkan biaya yang relatif banyak
- Beberapa perusahaan yang menyediakan AIaaS.
    - Amazon Web Services (AWS)
    - Microsft Azure
    - Google Cloud Platform (GCP)
    - IBM
    - OpenAI
- Tipe-tipe AIaaS
    - Bots/Chatbots
        - Amazon lex (AWS), Azure Bot Service (Microsft Azure), DialogFlow(GCP)
    - APIs 
        - Amazon Rekognition, Amazon Comprehend
        - Azure Cognitive Service, Azure Speech Services
        - Google Cloud Vision, Cloud Natural language
    - Machine learning
        - Amazon SageMaker
        - Azure Machine learning
        - Google Cloud AI
- Keuntungan menggunakan AI as a Services adalah sebagai berikut.
    - Pengurangan Cost 
    - Low-mode
    - Proses deployment cepat
    - Flexibility
    - Usability
    - Scalability
    - Customization
- Kelemahan menggunakan AI as a Services adalah sebagai berikut.
    - Cost yang berlebih dalam jangka panjang
    - Privasi data
    - Keamanan
    - Transparasi
    - Vendor lock-in

### OpenAI API
- Target :
    - Mendesain prompt yang tepat (prompt engineering)
    - Membuat aplikasi berbasis AI menggunakan API OpenAI
- [Dokumentasi OpenAI API](https://platform.openai.com/docs/introduction)
- [OpenAI API Reference](https://platform.openai.com/docs/api-reference/introduction)
- Opsional : install library [Go OpenAI](https://github.com/sashabaranov/go-openai)
- OpenAI API Pricing
    - Start for free , mendapatkan free credit senilai $5 yang dapat digunakan selama 3 bulan pertama semenjak register
    - Pay as you go, setiap model memiliki tarif yang berbeda berbeda
    - Lihat detail pada [Link](https://openai.com/pricing)
- Langkah mendapatkan API key untuk mengakses OpenAI API
    - Buka situs web [OpenAI](https://platform.openai.com/docs/api-reference)
    - Lakukan Sign Up untuk memiliki account openAI atau Login
    - Masuk ke bagian API keys atau melalui halaman berikut [OpenAI API Key](https://platform.openai.com/account/api-keys)
    - Setelah masuk ke halaman API key klik tombol create new secret key