#program untuk menghitung luas dan keliling lingkaran

phi = 3,14
pilih = 0

while pilih !=3:
    print("Program untuk menghitung luas dan keliling lingkarang")
    r = int(input("Masukan jari-jari lingkaran dalam satuan centi meter :"))
    print()
    
    #rumus
    luas = phi*r**2
    keliling = 2*phi*r
    
    print("Silahkan pilih nomor program :\n1. Luas lingkaran\n2. Keliling lingkaran\n3. Keluar dari program")
    pilih = int(input("Masukan nomor :"))
    
    
    if pilih==1:
        print("Luas lingkaran dengan jari-jari",r,"cm","adalah",luas,"cm\n")
        
    elif pilih==2:
        print("keliling lingkaran dengan jari-jari",r,"cm","adalah",keliling,"cm\n")
        
    elif pilih==3:
        print("Terima kasih telah menggunakan program ini")
        
    else:
        print("angka yang anda masukan salah\n")