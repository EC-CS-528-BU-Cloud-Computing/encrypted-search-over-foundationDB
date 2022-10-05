package edu.bu.fdb;

public class ServiceTest {
    public static void main(String[] args) {
        FDBService.UploadAll("E:\\BU\\Semester2\\EC504\\review");
        FDBService.DownloadAFile("Final Review.pdf");
        FDBService.ClearAll("E:\\BU\\Semester2\\EC504\\review");
    }
}
