# Summary

## (25) Algoritma AI

### Tipe - tipe Artificial intelligence (AI)
- Berdasarkan level kecerdasan :
    - narrow / weak AI
        - AI untuk melaksanakan tugas yang spesifik
        - Kemampuannya sangat dioptimaliasi untuk sebuah tugas dan dapat melampaui kemampuan manusia
        - Contoh : facial recognition, natural language processing, autonomous vehicle, dan voice assistants
    - general / strong AI
    - super AI
- Berdasarkan Fungsionalitas :
    - reactive machine
        - Reactive machine : tipe AI yang melakukan reaksi terhadap input tanpa pengetahuan masa lampau dan hanya spesifik melakukan tugas tertentu
        - Contoh : Machine learning, recomendation engine, deep blue (Mesin catur IBM)
    - limited memory
        - Limited memory : pengembangan dari reactive machine, sudah dilengkapi dengan kemampuan mengevaluasi situasi masa lampau
        - Contoh : deep learning , reinforcement learning, self-drivinf cars
    - theory of mind
    - self - awarness

### Algoritma Artificial intelligence (AI)
- Definisi Algoritma adaalah sebuah set untruksi yang diberikan kepada komputer unutuk memecahkan masalah atau mencapai sebuah tujuan
- Dalam konteks AI algoritma digunakan agar komputer mampu melakukan kalkulasi dan pemrosesan data, melatih model machine learning dan deep learning, dan membuat prediksi berdasarkan data yang diberikan

### Machine Learning
- Subset dari AI yang fokus pada pengembangan algoritma dan model
- Memiliki kemampuan untuk belajar dan mengambil keputusan atau prediksi tanpa perlu diprogram secara eksplisit
- Kualitas dari model machine learning tergntung dari bagus/tidaknya data yang diberikan
-  Machine Learning taxonomy.
    - Supervised Learning
        - Model machine learning di training pada data berlabel, dimana nilai input - output sudah diketahui
        - Algoritma machine learning mempelajari cara me-mapping sebuah function dari input untuk menghasilkan output, Y = f(X)
        - Ada dua tipe supervised learning : regression dan classification
    - Unsupervised Learning
        - Model machine learning di training pada data tidak berlabel
        - Mencari pola, mengelompokan data berdasarkan kesamaan atribut pada dataset yang diberikan dan membuat keputusan berdasarkan apa yang ditemukan
        - Ada beberapa tipe supervised learning : clustering, anomaly detection, association rule learning, dimensionality reduction
    - Semi-supervised Learning
        - Perpaduan antara supervised learning dengan semi-supervised learning
        - Mengombinasikan data berlabel dan tidak berlabel untuk meningkatkan akurasi model machine learning
        - Digunakan saat sejumlah besar data yang tersedia tidak berlabel, namun diperlukan waktu yang lama atau biaya yang mahal untuk memberi label
        - Ada beberapa tipe supervised learning : Support vector machine (SVM), K-nearest neighbors, Decision trees, Label propagation, Label spreading, Multi-view clustering, Multi-view dimensionality reduction, dan Multi-view layering
    - Reinforcement learning
        -  Sebuah proses iteratif dimana agent berinteraksi dengan environment dan belajar daru feedback yang didapat
        - Belajar dengan cara memaksimalkan reward dan meminimalkan punishment dalam rangka mencapai goal yang diberikan
- Deep Learning
    -  Subset dari machine learning
    - Menggunakan artificial neural network untuk meniru proses pembelajaran manusia
    - Dibandingkan machine learning, dapat membangun korelasi antara data dengan lebih kompleks, namun membutuhkan dataset yang relatif lebih besar
    - Idealnya menggunakan GPU (graphic processing unit) dalam melatih model deep learning

### Frameworks dan library untuk Artificial intelligence (AI)

Python : 
- [Pandas](https://pandas.pydata.org/)
- [TensorFlow](https://www.tensorflow.org/)
- [Keras](https://keras.io/)
- [NumPy](https://numpy.org)
- [Scikit-learn](https://scikit-learn.org/)
- [PyTorch](https://pytorch.org/)

Golang :
- [golearn](https://github.com/golang-basic/golearn)
- [gorgonia](https://github.com/gorgonia/gorgonia)
- [eaopt](https://github.com/MaxHalford/eaopt)
- [gonum](https://github.com/gonum/gonum)
- [gomind](https://github.com/surenderthakran/gomind)