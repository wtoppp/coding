package main

import (
	"archive/tar"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Tarfile(src string, dst string) interface{} {
	s, _ := os.Open(src)
	defer s.Close()

	d, _ := os.Create(dst)
	defer d.Close()

	tw := tar.NewWriter(d)
	defer tw.Close()

	srcinfo, _ := os.Stat(src)
	hdr, _ := tar.FileInfoHeader(srcinfo, "")
	fmt.Println(hdr)
	tw.WriteHeader(hdr)

	_, err := io.Copy(tw, s)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return true

}
func Untarfile(tarsrc string, untarpath string) interface{} {
	tarfile, _ := os.Open(tarsrc)
	defer tarfile.Close()

	tr := tar.NewReader(tarfile)
	path, _ := os.Stat(untarpath)

	if path != nil {
		os.RemoveAll(untarpath)
	}
	os.Mkdir(untarpath, 0755)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		filename := untarpath + hdr.Name
		fw, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0777)
		defer fw.Close()
		_, err = io.Copy(fw, tr)
		if err != nil {
			return err
		}
	}
	return true
}
func Copyfile1(srcpath string, dstpath string) {
	sinfo, err := os.Stat(srcpath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("sinfo", sinfo.Name())

	rh, _ := os.Open(srcpath)
	defer rh.Close()

	fname := filepath.Base(srcpath)
	dstfile := dstpath + "/" + fname

	wh, _ := os.Create(dstfile)
	defer wh.Close()

	io.Copy(wh, rh)
	return
}
func Copyfile2(srcpath string, dstpath string) {
	sinfo, err := os.Stat(srcpath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("sinfo", sinfo.Name())

	rh, _ := os.Open(srcpath)
	defer rh.Close()

	fname := filepath.Base(srcpath)
	dstfile := dstpath + "/" + fname

	wh, _ := os.Create(dstfile)
	defer wh.Close()

	//����һ��������buf,��Դ�ļ��ж�ȡһ����������(�Ϊlen(buf)),nΪ���������ݳ���,д�뵽buf��
	buf := make([]byte, 1024)
	for {
		n, err := rh.Read(buf)
		fmt.Println("nnnnnn", n)
		if err != nil && err == io.EOF {
			fmt.Println(err)
		}

		if n == 0 {
			break
		}
		//�Ѷ�ȡ��������,�浽һ��������,��д��Ŀ���ļ���,�Ӷ���ɸ����ļ�������
		tmp := buf[:n]
		if _, err = wh.Write(tmp); err != nil {
			fmt.Println(err)
		}
	}

	return
}
func Copyfile3(srcpath string, dstpath string) {
	sinfo, err := os.Stat(srcpath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("sinfo", sinfo.Name())

	input, _ := ioutil.ReadFile(srcpath)

	fname := filepath.Base(srcpath)
	dstfile := dstpath + "/" + fname

	ioutil.WriteFile(dstfile, input, 0644)

	return
}

func main() {
	src := "log1.txt"
	dst := "ok3.tar"
	Tarfile(src, dst)

	tarsrc := "ok3.tar"
	untarpath := "/tmp/test1/"
	Untarfile(tarsrc, untarpath)

	topdir := "/root"
	fh, _ := os.OpenFile(topdir, os.O_RDONLY, os.ModeDir)
	defer fh.Close()
	finfo, _ := fh.Readdir(-1)
	fmt.Printf("finfo %T\n", finfo)
	for _, val := range finfo {
		if !val.IsDir() {
			fname := val.Name()
			if strings.HasSuffix(fname, ".txt") {
				fmt.Println(".txt�ļ�", fname)
				abspath := topdir + "/" + fname
				fmt.Println(abspath)
				//����3�ָ����ļ��ķ���������,���ڴ��ļ����鲻��ioutil�ķ���
				Copyfile1(abspath, "/tmp/test1")
				//Copyfile2(abspath,"/tmp/test1")
				//Copyfile3(abspath,"/tmp/test1")
			}
		}
	}
}