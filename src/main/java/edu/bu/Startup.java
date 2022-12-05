package edu.bu;

import edu.bu.search.ClusionDlsD;
import edu.bu.util.Analysis;

import java.io.BufferedReader;
import java.io.File;
import java.io.InputStreamReader;
import java.security.Security;
import java.util.List;

public class Startup {
    public static void main(String[] args) {
        Security.addProvider(new org.bouncycastle.jce.provider.BouncyCastleProvider());
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
        while (true) {
            try {
                System.out.print("Enter operation (upload / download / clear / search): ");
                String op = reader.readLine();
                ClusionDlsD dlsd = new ClusionDlsD();
                if("upload".equals(op)) {
                    System.out.print("Enter path: ");
                    String path = reader.readLine();
                    dlsd.generateDlsD(path);
                    long filesSize = dlsd.computeFilesSize(new File(path));
                    Analysis.addFilesSize(filesSize);
                    dlsd.uploadDlsD();
                }
                if("download".equals(op)) {
                    dlsd.downloadDlsD();
                    Analysis.computeRate();
                }
                if("clear".equals(op)) {
                    dlsd.clearDlsD();
                }
                if("search".equals(op)) {
                    while (true) {
                        System.out.println("Enter keyword (finish by enter quit):");
                        String keyword = reader.readLine();

                        if ("quit".equals(keyword)) {
                            System.out.println("end.");
                            break;
                        }

                        List<String> strings = dlsd.queryToken(keyword);
                        System.out.println(strings);
                    }
                }
                System.out.print("exit?(Y/N)");
                String check = reader.readLine();
                if("Y".equals(check)) {
                    break;
                }
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
    }
}
