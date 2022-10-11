package edu.bu.fdb;

public class ServiceTest {
    public static void main(String[] args) {
        FDBService.UploadAll(".\\src\\test\\files");
        FDBService.DownloadAFile("Final Review.pdf");
        FDBService.ClearAll(".\\src\\test\\files");
    }
}
